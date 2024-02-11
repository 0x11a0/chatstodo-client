export function LinkButton({ title, onClick, isActivated, children }) {
    const className = `
                min-w-24
                px-4 rounded-full
                bg-dark-background-tertiary ` + (isActivated ? "text-dark-font-primary" : "text-dark-font-tertiary");
    return (
        <button onClick={onClick}
            className={className}
        >{title}</button>
    );
}

export function ListEntry({ children }) {
    return (
        <label className="
            flex-1
            w-full flex flex-row justify-between
            border-b border-dark-border-primary
            py-8 px-4 mb-8
            ">
            {children}
        </label>
    );
}
