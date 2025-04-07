import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import path from "node:path";
import tailwindcss from "@tailwindcss/vite";
import svgr from "vite-plugin-svgr";

const alias = (x: string): string => path.resolve(__dirname, x);

export default defineConfig({
	plugins: [react(), tailwindcss(), svgr()],
	resolve: {
		alias: {
			"@": alias("./src"),
		},
	},
});
