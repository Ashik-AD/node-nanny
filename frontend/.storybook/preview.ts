import type { Preview } from '@storybook/react'
import {
    withRouter,
    reactRouterParameters,
} from 'storybook-addon-remix-react-router'
import '../src/style.css'
const preview: Preview = {
    decorators: [withRouter],
    parameters: {
        controls: {
            matchers: {
                color: /(background|color)$/i,
                date: /Date$/i,
            },
        },
    },
}

export default preview
