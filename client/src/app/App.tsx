import type { FC } from "react";
import "./styles/global.css";
import "@/shared/fonts";

import { Icon } from "@/shared/ui";

const App: FC = () => {
	return (
		<div className="w-5 h-5 text-accent-primary">
			<Icon icon="screencast" />
		</div>
	);
};

export { App };
