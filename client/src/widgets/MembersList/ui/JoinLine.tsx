import { Icon, Title } from "@/shared/ui";
import type { FC } from "react";
import { memo } from "react";

const JoinLineComponent: FC = () => {
	return (
		<div className="flex items-center w-full gap-x-2 cursor-pointer box-border py-1 px-2 hover:bg-bg-tertiary rounded-lg">
			<div className="rounded-full h-10 w-10 flex">
				<div className="h-9.5 w-9.5 flex m-auto bg-bg-quaternary rounded-full">
					<div className="h-5 w-5 m-auto text-primary">
						<Icon icon="arrow_up" />
					</div>
				</div>
			</div>
			<Title>Join room</Title>
		</div>
	);
};

export const JoinLine = memo(JoinLineComponent);
