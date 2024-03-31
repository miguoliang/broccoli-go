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
    <ul className={"flex space-x-2"}>
      <li>
        <button
          className={"p-2 rounded-md bg-blue-500 text-white"}
          onClick={login}
        >
          Login
        </button>
      </li>
      <li>
        <a href={signUpUrl} className={"p-2 rounded-md bg-blue-500 text-white"}>
          Register
        </a>
      </li>
    </ul>
  );
}
