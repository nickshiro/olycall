import { Suspense } from "react";
import type { FC, ReactNode } from "react";

export interface MainProps {
	children: ReactNode;
}

const Main: FC<MainProps> = ({ children }) => {
	return <Suspense>{children}</Suspense>;
};

export { Main };
