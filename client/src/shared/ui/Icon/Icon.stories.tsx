import type { Meta, StoryObj } from "@storybook/react";

import { Icon } from "./Icon";

const meta: Meta<typeof Icon> = {
	component: Icon,
	title: "Icon",
};

export default meta;
type Story = StoryObj<typeof Icon>;

export const Default: Story = {
	args: {
		icon: "screencast",
	},
	decorators: [
		(Story) => (
			<div className="w-25 h-25 text-primary">
				<Story />
			</div>
		),
	],
};
