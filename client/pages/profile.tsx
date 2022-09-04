import { useState, useEffect, useCallback } from "react";
import { supabase } from "../utils/supabaseClient";
import { useRouter } from "next/router";
import { User } from "@supabase/supabase-js";
import { GetServerSideProps } from "next";

type Props = {
  user: User;
};

export default function Profile(props: Props) {
  const { user } = props;
  const router = useRouter();

  async function signOut() {
    await supabase.auth.signOut();
    router.push("/sign-in");
  }

  return (
    <div style={{ maxWidth: "420px", margin: "96px auto" }}>
      <h2>Hello, {user.email}</h2>
      <p>User ID: {user.id}</p>
      <button onClick={signOut}>Sign Out</button>
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
