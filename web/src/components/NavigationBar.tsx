import { Link } from "react-router-dom";

export default function NavigationBar() {
  return (
    <nav>
      <ul className={"grid grid-cols-3 text-center"}>
        <li>
          <Link to={"/"}>Home</Link>
        </li>
        <li>
          <a href="https://miguoliang.github.io/broccoli-go/">Docs</a>
        </li>
        <li>
          <Link to={"/prices"}>Prices</Link>
        </li>
      </ul>
    </nav>
  );
}
