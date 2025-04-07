import { memo } from "react";
import type { FC } from "react";

import { Button } from "@/shared/ui";

const JoinButton: FC = memo(() => {
	return (
		<Button className="bg-bg-tertiary hover:bg-bg-quaternary">
			<p className="text-accent-primary text-base">Join room</p>
		</Button>
	);
});

export { JoinButton };
