import { useCallback } from "react";
import userManager, { signUpUrl } from "../oidc.ts";

export default function AnonymousActions() {
  const login = useCallback(() => {
    userManager
      .signinRedirect()
      .then(() => {
        console.log("Redirecting to login");
      })
      .catch((error) => {
        console.error("Error redirecting to login", error);
      });
  }, []);

  return (
    <ul className={"flex space-x-10 ml-auto"}>
      <li>
        <button className="border-0 underline underline-offset-8 hover:bg-white hover:no-underline" onClick={login}>
          Login
        </button>
      </li>
      <li>
        <a className="underline underline-offset-8 hover:no-underline" href={signUpUrl}>
          Register
        </a>
      </li>
    </ul>
  );
}
