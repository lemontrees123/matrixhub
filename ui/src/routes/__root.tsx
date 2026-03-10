import {
  createRootRoute, Outlet, HeadContent,
} from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'

import i18n from '@/i18n'

export const Route = createRootRoute({
  component: () => (
    <>
      <HeadContent />
      <Outlet />
      <TanStackRouterDevtools initialIsOpen={false} />
    </>
  ),
  head: () => ({
    meta: [{
      title: i18n.t('translation.title'),
    }],
    links: [
      {
        rel: 'icon',
        href: '/favicon.ico?',
      },
    ],
  }),
})
