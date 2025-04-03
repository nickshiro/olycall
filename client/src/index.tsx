import { createRoot } from "react-dom/client";
import App from "@/app/App";

// biome-ignore lint/style/noNonNullAssertion: <explanation>
createRoot(document.getElementById("root")!).render(<App />);
