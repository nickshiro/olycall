import type { Meta, StoryObj } from "@storybook/react";

import { Panel } from "./Panel";

const meta: Meta<typeof Panel> = {
	component: Panel,
	title: "Widgets/Panel",
};

export default meta;

type Story = StoryObj<typeof Panel>;

export const Default: Story = {
	args: {},
	decorators: [
		(Story) => (
			<div className="w-89.5">
				<Story />
			</div>
		),
	],
};
