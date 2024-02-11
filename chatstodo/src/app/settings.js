import { ListEntry, LinkButton } from "./components";
import Switch from "react-switch";
import { useContext } from "react";
import { ThemeContext } from "./context";

export function SettingsPanel() {
    const { theme, setTheme } = useContext(ThemeContext);

    return (
        <div className="
            flex-1
            flex flex-col
            justify-start items-center
            p-4 w-full
        ">
            <SettingsSegment title="General">
                <ListEntry>
                    <span>Dark Mode</span>
                    <Switch onChange={() => {
                        setTheme(theme === "dark" ? "light" : "dark")
                    }}

                        checked={theme === "dark"} />
                </ListEntry>

                <ListEntry>
                    <span>Telegram</span>
                    <LinkButton title="Linked" isActivated />
                </ListEntry>

            </SettingsSegment>

            <SettingsSegment title="Tasks">
                <ListEntry>
                    <span>Google Tasks</span>
                    <LinkButton title="Link Now" />
                </ListEntry>
            </SettingsSegment>

            <SettingsSegment title="Events">
                <ListEntry>
                    <span>Google Calendar</span>
                    <LinkButton title="Link Now" />
                </ListEntry>
            </SettingsSegment>
        </div>

    );

}


function SettingsSegment({ title, children }) {
    return (
        <div className="
            flex-1 flex flex-col
            w-full m-2
        ">
            <h3 className="
                text-dark-font-primary 
                font-primary font-bold
                text-m
            ">{title}</h3>

            {children}
        </div>
    );
}
