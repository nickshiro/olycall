import { Avatar, Button, Icon, Title } from "@/shared/ui";
import type { FC } from "react";
import { memo } from "react";

const PanelComponent: FC = () => {
	return (
		<div className="bg-bg-secondary p-2 box-border rounded-2xl flex flex-col gap-y-4">
			<div className="flex w-full justify-between items-center">
				<section className="flex items-center gap-x-2">
					<Avatar src="//m.dedkov.space/meme/dd" isActive={true} />
					<Title>Pavel Durov</Title>
				</section>
				<button
					type="button"
					className="h-8 w-8 rounded-lg text-primary flex hover:bg-bg-tertiary cursor-pointer"
				>
					<div className="w-5 h-5 m-auto">
						<Icon icon="settings" />
					</div>
				</button>
			</div>
			<div className="flex items-center gap-x-2 h-8 box-border">
				<Button>
					<Icon icon="video" />
				</Button>
				<Button>
					<Icon icon="screencast" />
				</Button>
				<Button>
					<Icon icon="microphone_muted" />
				</Button>
				<Button>
					<Icon icon="headphones" />
				</Button>
				<Button>
					<Icon icon="leave" />
				</Button>
			</div>
		</div>
	);
};

export const Panel = memo(PanelComponent);
