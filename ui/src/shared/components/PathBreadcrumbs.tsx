import {
  Anchor,
  Breadcrumbs,
  Text,
} from '@mantine/core'

export interface ModelPathBreadcrumbsProps {
  name: string
  treePath?: string
  onPathClick?: (path: string) => void
}

export function PathBreadcrumbs({
  name,
  treePath,
  onPathClick,
}: ModelPathBreadcrumbsProps) {
  const pathSegments = treePath?.split('/').filter(Boolean) ?? []

  return (
    <Breadcrumbs separator="/" separatorMargin="xs">
      <Anchor
        c="gray.9"
        onClick={(event) => {
          event.preventDefault()
          onPathClick?.('')
        }}
      >
        {name}
      </Anchor>

      {pathSegments.map((segment, index) => {
        const pathToSegment = pathSegments.slice(0, index + 1).join('/')

        if (!onPathClick || index === pathSegments.length - 1) {
          return (
            <Text key={pathToSegment} c="gray.9">
              {segment}
            </Text>
          )
        }

        return (
          <Anchor
            key={pathToSegment}
            c="gray.9"
            onClick={(event) => {
              event.preventDefault()
              onPathClick(pathToSegment)
            }}
          >
            {segment}
          </Anchor>
        )
      })}
    </Breadcrumbs>
  )
}
