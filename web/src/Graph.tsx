import {
  Button,
  Description,
  Dialog,
  DialogPanel,
  DialogTitle,
  Transition,
  TransitionChild,
} from "@headlessui/react";
import { ErrorMessage, Field, FieldArray, Form, Formik } from "formik";
import {
  PropsWithChildren,
  createContext,
  useContext,
  useMemo,
  useState,
} from "react";
import { MdDelete } from "react-icons/md";
import ReactFlow, {
  Background,
  BackgroundVariant,
  Controls,
  MiniMap,
  Panel,
  ReactFlowProvider,
} from "reactflow";
import * as Yup from "yup";
import SearchBox from "./components/SearchBox";
import useGraphStore from "./stores/GraphStore";

const AddNodeDialogContext = createContext<
  PropsWithChildren<{
    isOpen: boolean;
    onClose: () => void;
  }>
>({ isOpen: false, onClose: () => {} });

const Graph = () => {
  const { nodes, edges } = useGraphStore();
  const [isOpen, setIsOpen] = useState(false);
  const dialogContextValue = useMemo(
    () => ({ isOpen, onClose: () => setIsOpen(false) }),
    [isOpen],
  );
  return (
    <AddNodeDialogContext.Provider value={dialogContextValue}>
      <ReactFlowProvider>
        <ReactFlow nodes={nodes} edges={edges}>
          <Panel position="top-left">
            <div className="flex justify-between text-sm shadow-around bg-white py-1">
              <SearchBox />
              <button className="px-2" onClick={() => setIsOpen(true)}>
                Add Node
              </button>
            </div>
          </Panel>
          <Controls className="bg-white" />
          <MiniMap />
          <Background variant={BackgroundVariant.Dots} />
        </ReactFlow>
      </ReactFlowProvider>
      <AddNodeDialog />
    </AddNodeDialogContext.Provider>
  );
};

const AddNodeDialog = () => {
  const { isOpen, onClose } = useContext(AddNodeDialogContext);
  return (
    <Transition appear show={isOpen} as="div">
      <Dialog
        as="div"
        className="relative z-10 focus:outline-none"
        onClose={onClose}
      >
        {/* The backdrop, rendered as a fixed sibling to the panel container */}
        <TransitionChild
          as="div"
          className="fixed inset-0 bg-black/10"
          aria-hidden="true"
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        />
        {/* Full-screen container to center the panel */}
        <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
          {/* The actual dialog panel  */}
          <TransitionChild
            as={DialogPanel}
            className="max-w-xl space-y-4 bg-white p-12 rounded-md shadow-around"
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <DialogTitle className="font-bold">Deactivate account</DialogTitle>
            <Description>
              This will permanently deactivate your account
            </Description>
            <AddNodeForm />
          </TransitionChild>
        </div>
      </Dialog>
    </Transition>
  );
};

type AddNodeFormValues = {
  name: string;
  type: string;
  props: Array<{ key: string; value: string }>;
};

const validationSchema = Yup.object({
  name: Yup.string().required("Name is required"),
  type: Yup.string().required("Type is required"),
});

const initialValues: AddNodeFormValues = {
  name: "",
  type: "",
  props: [],
};

const AddNodeForm = () => {
  const handleSubmit = (values: AddNodeFormValues) => {
    // Handle form submission
    console.log(values);
  };
  const { onClose } = useContext(AddNodeDialogContext);
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={handleSubmit}
    >
      {({ values }) => (
        <Form className="flex flex-col gap-4">
          <Field
            type="text"
            id="name"
            name="name"
            className="w-full border border-gray-300 rounded-sm p-1"
            placeholder="Name"
          />
          <ErrorMessage
            name="name"
            component="div"
            className="text-red-500 text-sm"
          />
          <Field
            type="text"
            id="type"
            name="type"
            className="w-full border border-gray-300 rounded-sm p-1"
            placeholder="Type"
          />
          <ErrorMessage
            name="type"
            component="div"
            className="text-red-500 text-sm"
          />
          <PropArrayField {...values} />
          {/* Add more prop fields as needed */}
          <Button
            className="rounded-sm bg-black py-1.5 px-3 text-sm/6 font-semibold text-white shadow-inner shadow-white/10 focus:outline-none mt-5"
            type="submit"
          >
            Save
          </Button>
          <Button
            className="rounded-sm bg-white py-1.5 px-3 text-sm/6 font-semibold focus:outline-none border"
            type="reset"
            onClick={onClose}
          >
            Close
          </Button>
        </Form>
      )}
    </Formik>
  );
};

const PropArrayField = ({ props }: AddNodeFormValues) => {
  return (
    <FieldArray name="props">
      {({ push, remove }) => (
        <div className="flex flex-col gap-2">
          <h3>Properties</h3>
          {props.map((_, index) => (
            <div key={index} className="flex gap-2 rounded-sm">
              <Field
                type="text"
                id={`props.${index}.key`}
                name={`props.${index}.key`}
                placeholder="Key"
                className="flex-grow-1 border border-gray-300 rounded-sm"
              />
              <ErrorMessage name={`props.${index}.key`} component="div" />
              <Field
                type="text"
                id={`props.${index}.value`}
                name={`props.${index}.value`}
                placeholder="Value"
                className="flex-grow-1 border border-gray-300 rounded-sm"
              />
              <ErrorMessage name={`props.${index}.value`} component="div" />
              <Button
                className="flex-grow-0 hover:text-red-800"
                type="button"
                onClick={() => remove(index)}
              >
                <MdDelete />
              </Button>
            </div>
          ))}
          <Button
            type="button"
            className="px-1 underline underline-offset-8 hover:no-underline"
            onClick={() => {
              if (!props.find(({ key }) => key === "")) {
                push({ key: "", value: "" });
              }
            }}
          >
            Add Prop
          </Button>
        </div>
      )}
    </FieldArray>
  );
};

export default Graph;
