import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import path from "node:path";
import tailwindcss from "@tailwindcss/vite";

const alias = (x: string): string => path.resolve(__dirname, x);

export default defineConfig({
	plugins: [react(), tailwindcss()],
	resolve: {
		alias: {
			"@app": alias("./src/app"),
			"@entities": alias("./src/entities"),
			"@widgets": alias("./src/widgets"),
			"@features": alias("./src/features"),
			"@shared": alias("./src/shared"),
			"@pages": alias("./src/pages"),
		},
	},
});
