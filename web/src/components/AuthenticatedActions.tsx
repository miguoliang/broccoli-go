import { User } from "oidc-client-ts";
import { signOutUrl } from "../oidc.ts";

export default function AuthenticatedActions({ user }: Readonly<{ user: User }>) {
  return (
    <ul className={"flex space-x-2"}>
      <li>
        <a
          href={"/profile"}
          className={"inline-block p-2 rounded-md bg-blue-500 text-white"}
        >
          {user?.profile.email}
        </a>
      </li>
      <li>
        <a
          href={signOutUrl}
          className={"inline-block p-2 rounded-md bg-blue-500 text-white"}
        >
          Logout
        </a>
      </li>
    </ul>
  );
}
