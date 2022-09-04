import { Provider } from "@supabase/supabase-js";
import { NextPage } from "next";
import { supabase } from "../utils/supabaseClient";

const SignIn: NextPage = () => {
  const signIn = async (provider: Provider) => {
    await supabase.auth.signIn({ provider });
  };

  const buttonStyles = "p-2 rounded-md border-gray-500 border";

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="flex flex-col gap-2">
        <button
          className={buttonStyles}
          onClick={async () => await signIn("google")}
        >
          Sign In With Google
        </button>
      </div>
    </div>
  );
};

export default SignIn;
