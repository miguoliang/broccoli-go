import { useGetPLink } from "./api.ts";

export default function Profile() {

  const paymentLink = useGetPLink();

  return (
    <div>
      <h1>Profile</h1>
      {paymentLink.data && <a href={paymentLink.data.url}>Buy something!</a>}
    </div>
  );
}