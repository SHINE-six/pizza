import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    screens: {
      sm: "640px",
      md: "768px",
      lg: "1024px",
      xl: "1280px",
    },
    colors: {
      primary: {
        light: "#c9cfff",
        DEFAULT: "#8c8de6",
        dark: "#313c5e",
      },
      secondary: {
        light: "#FFD700",
        DEFAULT: "#FFD700",
        dark: "#FFD700",
      },
      tertiary: {   // mostly for navbar
        light: "#8c8de6",
        DEFAULT: "#de1f12",
        dark: "#313c5e",
      },
      white: "#FFFFFF",
      black: "#000000",
    },
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },
  },
  plugins: [],
};
export default config;
