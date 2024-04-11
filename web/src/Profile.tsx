import { useGetPLink, useGetProfileSubscriptions } from "./api.ts";

export default function Profile() {

  const paymentLink = useGetPLink();

  const subscriptions = useGetProfileSubscriptions();

  return (
    <div>
      <h1>Profile</h1>
      {paymentLink.data && <a href={paymentLink.data.url}>Buy something!</a>}
      {subscriptions.data && <div>
        <h2>Subscriptions</h2>
        <ul>
          {subscriptions.data.map(sub => (
            <li key={sub.id}>{sub.interval_count} {sub.interval}</li>
          ))}
        </ul>
      </div>}
    </div>
  );
}