import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute(
  '/_layout/(app)/projects/$projectId/models/',
)({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/_layout/(app)/projects/$projectId/models/"!</div>
}
