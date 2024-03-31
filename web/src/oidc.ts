import {
  User,
  UserManager,
  UserManagerSettings,
  WebStorageStateStore,
} from "oidc-client-ts";
import { create } from "zustand";

const clientId = "5p99s5nl7nha5tfnpik3r0rb7j";
const scope = "openid profile aws.cognito.signin.user.admin";
const authority =
  "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_Qbzi9lvVB";

// Create cognito sign-up & sign-out url
const query = new URLSearchParams();
query.append("client_id", clientId);
query.append("response_type", "code");
query.append("scope", scope);
query.append("redirect_uri", window.location.origin);

export const signUpUrl = `https://broccoli-go-user-pool-domain.auth.us-east-1.amazoncognito.com/signup?${query.toString()}`;
export const signOutUrl = signUpUrl.replace(/signup/, "logout");

const config: UserManagerSettings = {
  authority, // Replace with actual endpoint
  client_id: clientId, // Replace with actual client ID
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
