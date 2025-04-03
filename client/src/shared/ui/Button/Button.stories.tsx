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
		isActive: false,
		onClick: () => {},
		children: <Icon icon="screencast" />,
	},
	decorators: [
		(Story) => (
			<div className="box-border" style={{ width: 62 }}>
				<Story />
			</div>
		),
	],
};
