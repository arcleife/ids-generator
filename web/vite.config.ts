import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
    base: "/ids-gen",
    plugins: [react()],
    preview: {
        port: 4001,
        strictPort: true,
    },
    server: {
        port: 4001,
        strictPort: true,
        host: true,
        origin: "http://0.0.0.0:4001",
    },
});
