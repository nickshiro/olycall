import type { Meta, StoryObj } from "@storybook/react";

import { ToggleButton } from "./ToggleButton";
import { Icon } from "../Icon";

const meta: Meta<typeof ToggleButton> = {
	component: ToggleButton,
	title: "Shared/ToggleButton",
};

export default meta;

type Story = StoryObj<typeof ToggleButton>;

export const Default: Story = {
	args: {
		children: (
			<div className="h-5 w-5">
				<Icon icon="microphone_muted" />
			</div>
		),
		defaultToggle: true,
		onClick: (x) => {
			return !x;
		},
	},
	decorators: [
		(Story) => (
			<div className="flex gap-2">
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
