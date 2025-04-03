import type { Meta, StoryObj } from "@storybook/react";

import { Name } from "./Name";

const meta: Meta<typeof Name> = {
	component: Name,
	title: "Shared/Name",
};

export default meta;
type Story = StoryObj<typeof Name>;

export const Default: Story = {
	args: {
		children: "Pavel Durov",
	},
};
