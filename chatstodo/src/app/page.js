"use client"
import { FiRefreshCcw } from "react-icons/fi";
import { useContext, useEffect, useState } from "react";
import { ThemeContext } from "./context";
import { HomePanel, MobileHome } from "./home";
import { MobileSettings, SettingsPanel } from "./settings";
import { BotPanel, MobileBots } from "./bot";
import { FaRobot, FaSmileBeam } from "react-icons/fa";
import { IoMdSettings } from "react-icons/io";

const MOBILE_BREAKPOINT = 480;

export default function Home() {
    const [theme, setTheme] = useState("dark");

    useEffect(() => {
        setTheme(window.matchMedia('(prefers-color-scheme: dark)').matches ? "dark" : "light");
    }, []);
    console.log("theme: ", theme);

    const [tabIndex, setTabIndex] = useState(1);
    const [width, setWidth] = useState(0);

    useEffect(() => {
        function handleWindowSizeChange() {
            setWidth(window.innerWidth);
        }
        window.addEventListener('resize', handleWindowSizeChange);
        return () => {
            window.removeEventListener('resize', handleWindowSizeChange);
        }
    }, [setWidth]);

    const [hasData, setHasData] = useState(false);

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

                    <button className="
                        text-dark-font-primary 
                        flex-1 text-start
                        font-primary font-bold
                        text-2xl"
                        onClick={() => setTabIndex(1)}
                    >ChatsToDo</button>

                    {width <= MOBILE_BREAKPOINT &&
                        <>
                            <button onClick={() => setTabIndex(2)}>
                                <FaRobot className="
                                    text-2xl text-dark-font-primary mx-2"/>
                            </button>

                            <button onClick={() => setTabIndex(3)}>
                                <IoMdSettings className="
                                    text-2xl text-dark-font-primary mx-2"/>
                            </button>

                        </>
                    }
                    <button onClick={() => setHasData(true)}>
                        <FiRefreshCcw className="
                            text-2xl text-dark-font-primary mx-2"/>
                    </button>
                </div>
                {width > MOBILE_BREAKPOINT &&
                    <>
                        <div className="
                flex-1 flex flex-row
                justify-start items-start
                w-screen h-screen p-0
            ">

                            <div className="
                    flex-1 flex flex-col justify-start items-first
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

                    </>
                }
                {width <= MOBILE_BREAKPOINT &&
                    <>
                        {tabIndex == 1 &&
                            <MobileHome hasData={hasData} />
                        }
                        {tabIndex == 2 &&
                            <MobileBots />
                        }
                        {tabIndex == 3 &&
                            <MobileSettings />
                        }
                    </>
                }
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

