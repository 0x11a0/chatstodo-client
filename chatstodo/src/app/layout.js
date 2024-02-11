import "./globals.css";

export const metadata = {
    title: "ChatToDo",
    description: "Web client for ChatToDo application",
};

export default function RootLayout({ children }) {

    return (
        <html lang="en">
            <body className="h-screen w-screen">
                {children}
            </body>
        </html>
    );
}
