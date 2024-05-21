/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{html,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      boxShadow: {
        around: "0 0 2px 1px rgba(0,0,0,0.08)"
      },
    },
  },
  plugins: [],
};
