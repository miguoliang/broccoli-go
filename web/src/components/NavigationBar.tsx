import SearchBox from "./SearchBox.tsx";
import { Menu } from "@headlessui/react";
import { Link } from "react-router-dom";

const activeLinkClasses = "bg-blue-500 text-white p-2 rounded-md";

const defaultLinkClasses = "p-2 rounded-md";

export default function NavigationBar() {
  return (
    <nav>
      <ul className={"flex space-x-2 items-stretch justify-center"}>
        <li className={"mr-5"}>
          <SearchBox />
        </li>
        <li className={"flex items-center"}>
          <Menu as="div" className="relative inline-block text-left">
            <Menu.Button className={"py-2"}>Create</Menu.Button>
            <Menu.Items>
              <div
                className={
                  "absolute flex flex-col border rounded-md bg-white p-2 shadow-md"
                }
              >
                <Menu.Item>
                  {({ active }) => (
                    <Link
                      className={
                        active ? activeLinkClasses : defaultLinkClasses
                      }
                      to={"/create-vertex"}
                    >
                      Product
                    </Link>
                  )}
                </Menu.Item>
                <Menu.Item>
                  {({ active }) => (
                    <Link
                      className={
                        active ? activeLinkClasses : defaultLinkClasses
                      }
                      to={"/create-vertex"}
                    >
                      Manufacture
                    </Link>
                  )}
                </Menu.Item>
              </div>
            </Menu.Items>
          </Menu>
        </li>
      </ul>
    </nav>
  );
}
