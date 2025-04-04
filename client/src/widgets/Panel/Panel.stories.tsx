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
			<div style={{ width: 358 }}>
				<Story />
			</div>
		),
	],
};
