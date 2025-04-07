import { BrowserRouter } from "react-router";
import type { FC } from "react";

import { Main as Router } from "./routers";
import { Main as Layout } from "./layouts";

import "./styles/global.css";
import "@/shared/fonts";

const App: FC = () => {
	return (
		<BrowserRouter>
			<Layout>
				<Router />
			</Layout>
		</BrowserRouter>
	);
};

export { App };
