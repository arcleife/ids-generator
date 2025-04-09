import "./App.css";

interface MyButtonProps {
    /** The text to display inside the button */
    title: string;
    /** Whether the button can be interacted with */
    disabled: boolean;
}

function MyButton({ title, disabled }: MyButtonProps) {
    return <button disabled={disabled}>{title}</button>;
}

function App() {
    return (
        <>
            <h1>IDs Generator</h1>
            <div className="card">
                <h2>Welcome to my app</h2>
                <MyButton title="I'm a button" disabled={true} />
            </div>
        </>
    );
}

export default App;
