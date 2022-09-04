import "../styles/globals.css";
import type { AppProps } from "next/app";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { supabase } from "../utils/supabaseClient";
import { AuthChangeEvent, Session } from "@supabase/supabase-js";

function MyApp({ Component, pageProps }: AppProps) {
  const router = useRouter();
  async function handleAuthChange(
    event: AuthChangeEvent,
    session: Session | null
  ) {
    await fetch("/api/auth", {
      method: "POST",
      headers: new Headers({ "Content-Type": "application/json" }),
      credentials: "same-origin",
      body: JSON.stringify({ event, session })
    });
  }

  useEffect(() => {
    const { data: authListener } = supabase.auth.onAuthStateChange(
      (event, session) => {
        handleAuthChange(event, session);
        if (event === "SIGNED_IN") {
          router.push("/profile");
        }
      }
    );

    return () => {
      authListener?.unsubscribe();
    };
  }, [router]);

  return <Component {...pageProps} />;
}

export default MyApp;
