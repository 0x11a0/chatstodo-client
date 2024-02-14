"use client"
import { FiRefreshCcw } from "react-icons/fi";
import { useContext, useEffect, useState } from "react";
import { ThemeContext } from "./context";
import { HomePanel } from "./home";
import { SettingsPanel } from "./settings";
import { BotPanel } from "./bot";

export default function Home() {
    const [theme, setTheme] = useState("dark");

    useEffect(() => {
        setTheme(window.matchMedia('(prefers-color-scheme: dark)').matches ? "dark" : "light");
    }, []);
    console.log("theme: ", theme);

    const [tabIndex, setTabIndex] = useState(1);

    return (
        <ThemeContext.Provider value={{ theme, setTheme }}>
            <main className="
                bg-dark-background-primary 
                flex flex-col items-center justify-start
                min-h-screen
                p-2"
            >

                <div className="
                    border-dark-border-primary
                    flex content-start justify-between items-center
                    w-screen px-3 h-24 border-b"
                >

                    <h2 className="
                        text-dark-font-primary 
                        flex-1
                        font-primary font-bold
                        text-2xl"

                    > ChatsToDo</h2>
                    <button>
                        <FiRefreshCcw className="
                    text-2xl text-dark-font-primary"/>
                    </button>
                </div>

                <div className="
                flex-1 flex flex-row
                justify-start items-start
                w-screen h-screen p-0
            ">

                    <div className="
                    flex flex-col justify-start items-first
                    bg-dark-background-secondary
                    w-32 min-w-32 max-w-32
                    h-screen p-2">

                        <Tab title="Home" index={1} setIndex={setTabIndex} />
                        <Tab title="Bots" index={2} setIndex={setTabIndex} />
                        <Tab title="Settings" index={3} setIndex={setTabIndex} />

                    </div>

                    {tabIndex == 1 &&
                        <HomePanel />
                    }
                    {tabIndex == 2 &&
                        <BotPanel />
                    }
                    {tabIndex == 3 &&
                        <SettingsPanel />
                    }
                </div>
            </main>
        </ThemeContext.Provider >
    );
}

function Tab({ title, index, setIndex }) {
    return (
        <button
            className="
                text-base text-dark-font-secondary font-primary
                border-b border-dark-border-primary text-start
                py-2"
            onClick={() => setIndex(index)}

        >{title}</button>
    );
}
