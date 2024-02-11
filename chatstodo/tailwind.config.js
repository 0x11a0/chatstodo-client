/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    ],
    theme: {
        extend: {
            backgroundImage: {
                "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
                "gradient-conic":
                    "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
            },
            colors: {

                "dark-background-primary": "#292929",
                "dark-background-secondary": "#313131",
                "dark-background-tertiary": "#494949",
                "dark-font-primary": "#FFC658",
                "dark-font-secondary": "#8D8D8D",
                "dark-font-tertiary": "#FFFFFF",
                "dark-border-primary": "#434343",

            }
        },
        fontFamily: {
            "primary": ["Roboto", "sans-serif"],
        }
    },
    darkMode: "class",
    plugins: [],
};
/*
    "light-background-primary": "#EDEDED",
    "light-background-secondary": "#FFFFFF",
    "light-font-primary": "#000000",
    "light-font-secondary": "#000000",
    "light-border-primary": "#C2C2C2",
    */
