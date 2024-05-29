import { Input } from "@headlessui/react";
import Downshift from "downshift";

export default function SearchBox() {
  const items = [
    { author: "Harper Lee", title: "To Kill a Mockingbird" },
    { author: "Lev Tolstoy", title: "War and Peace" },
    { author: "Fyodor Dostoyevsy", title: "The Idiot" },
    { author: "Oscar Wilde", title: "A Picture of Dorian Gray" },
    { author: "George Orwell", title: "1984" },
    { author: "Jane Austen", title: "Pride and Prejudice" },
    { author: "Marcus Aurelius", title: "Meditations" },
    { author: "Fyodor Dostoevsky", title: "The Brothers Karamazov" },
    { author: "Lev Tolstoy", title: "Anna Karenina" },
    { author: "Fyodor Dostoevsky", title: "Crime and Punishment" },
  ];

  return (
    <Downshift
      onChange={(selection) =>
        alert(
          selection
            ? `You selected "${selection.title}" by ${selection.author}. Great Choice!`
            : "Selection Cleared",
        )
      }
      itemToString={(item) => (item ? item.title : "")}
    >
      {({
        getInputProps,
        getItemProps,
        getMenuProps,
        inputValue,
        highlightedIndex,
        selectedItem,
        isOpen,
      }) => (
        <div>
          <div className="p-1 flex justify-between bg-transparent">
            <Input
              {...getInputProps()}
              className={"border-0 outline-0 flex-grow text-sm min-w-[200px] px-2 border-r placeholder:text-gray-300 bg-transparent"}
              placeholder={"Search"}
            />
          </div>
          <ul
            className={`absolute w-72 bg-white mt-1 border rounded-md shadow-md max-h-80 overflow-auto p-0 z-10 ${
              !(isOpen && items.length) && "hidden"
            }`}
            {...getMenuProps()}
          >
            {isOpen
              ? items
                  .filter(
                    (item) =>
                      !inputValue ||
                      item.title
                        .toLowerCase()
                        .includes(inputValue.toLowerCase()) ||
                      item.author
                        .toLowerCase()
                        .includes(inputValue.toLowerCase()),
                  )
                  .map((item, index) => (
                    <li
                      className={`py-2 px-3 shadow-sm flex flex-col ${highlightedIndex === index && "bg-blue-300"} ${selectedItem === item && "font-bold"} cursor-pointer`}
                      key={`${item.author}${index}`}
                      {...getItemProps({
                        item,
                        index,
                      })}
                    >
                      <span>{item.title}</span>
                      <span className="text-sm text-gray-700">
                        {item.author}
                      </span>
                    </li>
                  ))
              : null}
          </ul>
        </div>
      )}
    </Downshift>
  );
}
