import type { Meta, StoryObj } from "@storybook/react";

import { Button } from "./Button";
import { Icon } from "../Icon";

const meta: Meta<typeof Button> = {
	component: Button,
	title: "Shared/Button",
};

export default meta;

type Story = StoryObj<typeof Button>;

export const Default: Story = {
	args: {
		children: (
			<div className="h-5 w-5">
				<Icon icon="settings" />
			</div>
		),
	},
	decorators: [
		(Story) => (
			<div className="flex gap-2 text-accent-primary">
				<div className="w-25">
					<Story />
				</div>
				<div className="w-8">
					<Story />
				</div>
			</div>
		),
	],
};
