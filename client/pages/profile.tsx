import { supabase } from "../utils/supabaseClient";
import { useRouter } from "next/router";
import { Session, User } from "@supabase/supabase-js";
import { GetServerSideProps } from "next";
import { useMutation } from "@tanstack/react-query";

type Props = {
  user: User;
};

export default function Profile(props: Props) {
  const session = supabase.auth.session() as Session;
  console.log(session.access_token);

  const { isLoading, mutate } = useMutation(
    ["create-todo"],
    async (data: { description: string }) => {
      const res = await fetch("http://localhost:8080/todo", {
        method: "POST",
        body: JSON.stringify({
          description: data.description
        }),
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${session.access_token}`
        }
      });

      const json = await res.json();

      console.log(json);

      return json as {
        id: string;
        description: string;
        created_at: string;
        updated_at: string;
      };
    }
  );

  const { user } = props;
  const router = useRouter();

  async function signOut() {
    await supabase.auth.signOut();
    router.push("/sign-in");
  }

  return (
    <div className="min-h-screen flex flex-col items-center justify-center">
      <h2>Hello, {user.email}</h2>
      <p>User ID: {user.id}</p>
      <button onClick={signOut}>Sign Out</button>
      <button
        onClick={() =>
          mutate({
            description: `Todo: ${Math.round(Math.random() * 1000000)}`
          })
        }
      >
        {isLoading ? "Creating" : "Create Todo"}
      </button>
    </div>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req }) => {
  const { user } = await supabase.auth.api.getUserByCookie(req);

  if (!user) {
    return { props: {}, redirect: { destination: "/sign-in" } };
  }

  return { props: { user } };
};
