import { Projects } from '@matrixhub/api-ts/v1alpha1/project.pb'
import { queryOptions, useQuery } from '@tanstack/react-query'

// -- Query key factory --

export const projectKeys = {
  detail: (projectId: string) => ['projects', projectId] as const,
}

// -- Query options factory --

export function projectQueryOptions(projectId: string) {
  return queryOptions({
    queryKey: projectKeys.detail(projectId),
    queryFn: () => Projects.GetProject({ name: projectId }),
  })
}

// -- Custom hook --
export function useProject(projectId: string) {
  return useQuery(projectQueryOptions(projectId))
}
