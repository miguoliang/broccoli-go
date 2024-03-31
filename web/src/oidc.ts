import {
  User,
  UserManager,
  UserManagerSettings,
  WebStorageStateStore,
} from "oidc-client-ts";
import { create } from "zustand";

const clientId = "3e2lpi6o9804lrd46gbbubeoqt";
const scope = "openid profile aws.cognito.signin.user.admin";
const authority =
  "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_wqKTwXm3J";

// Create cognito sign-up & sign-out url
const query = new URLSearchParams();
query.append("client_id", clientId);
query.append("response_type", "code");
query.append("scope", scope);
query.append("redirect_uri", window.location.origin);

export const signUpUrl = `https://okra-test.auth.us-east-1.amazoncognito.com/signup?${query.toString()}`;
export const signOutUrl = signUpUrl.replace(/signup/, "logout");

const config: UserManagerSettings = {
  authority, // Replace with actual endpoint
  client_id: clientId, // Replace with actual client ID
  client_secret: "fo68ccbmsvfc37gf8d33h6cp79e54tjjgfent3mf0fehe7grlud", // Replace with actual client secret
  redirect_uri: "http://localhost:5173/", // Adjust based on your redirect logic
  revokeTokenTypes: ["refresh_token"],
  automaticSilentRenew: false,
  response_type: "code",
  scope, // Adjust scopes as needed
  userStore: new WebStorageStateStore({ store: window.localStorage }),
};

const userManager = new UserManager(config);

type AuthStore = {
  user: User | null;
  setUser: (user: User) => void;
};

export const useAuthStore = create<AuthStore>((set) => ({
  user: null,
  setUser: (user) => set({ user }),
}));

export default userManager;
