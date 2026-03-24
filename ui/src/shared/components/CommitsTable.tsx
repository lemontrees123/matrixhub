import {
  Anchor,
  Group,
  Text,
} from '@mantine/core'
import { useTranslation } from 'react-i18next'

import { CopyValueButton } from '@/shared/components/CopyValueButton'
import { DataTable } from '@/shared/components/DataTable'
import { formatDateTime } from '@/shared/utils/date'

import type { Commit } from '@matrixhub/api-ts/v1alpha1/model.pb'
import type { Pagination as PaginationData } from '@matrixhub/api-ts/v1alpha1/utils.pb'
import type { MRT_ColumnDef } from 'mantine-react-table'

interface CommitsTableProps {
  commits: Commit[]
  pagination?: PaginationData
  page: number
  loading?: boolean
  onPageChange: (page: number) => void
  onDetailClick?: (commit: Commit) => void
}

type CommitCellProps = Parameters<NonNullable<MRT_ColumnDef<Commit>['Cell']>>[0]

function CommitMessageCell({ row }: CommitCellProps) {
  return (
    <Text size="sm" py="xs">
      {row.original.message || '-'}
    </Text>
  )
}

function CommitHashCell({ row }: CommitCellProps) {
  const commitId = row.original.id

  if (!commitId) {
    return <Text size="sm">-</Text>
  }

  return (
    <Group gap="xs" wrap="nowrap">
      <Text size="sm" ff="monospace" truncate>
        {commitId}
      </Text>
      <CopyValueButton
        value={commitId}
        iconSize={12}
        color="gray-9"
      />
    </Group>
  )
}

function CommitAuthorDateCell({ row }: CommitCellProps) {
  const author = row.original.authorName || '-'
  const authorDate = formatDateTime(row.original.authorDate, 'YYYY-MM-DD HH:mm:ss')

  return (
    <Text size="sm" truncate>
      {`${author} ${authorDate}`}
    </Text>
  )
}

function CommitDetailCell({
  row,
  table,
}: CommitCellProps) {
  const { t } = useTranslation()
  const {
    onDetailClick,
  } = table.options.meta as {
    onDetailClick?: (commit: Commit) => void
  }
  const commitId = row.original.id

  if (!commitId) {
    return <Text size="sm">-</Text>
  }

  return (
    <Anchor
      size="sm"
      underline="never"
      onClick={() => onDetailClick?.(row.original)}
    >
      {t('shared.commitsTable.actions.detail')}
    </Anchor>
  )
}

export function CommitsTable({
  commits,
  pagination,
  page,
  loading,
  onPageChange,
  onDetailClick,
}: CommitsTableProps) {
  const { t } = useTranslation()

  const columns: MRT_ColumnDef<Commit>[] = [
    {
      id: 'message',
      header: t('shared.commitsTable.columns.message'),
      Cell: CommitMessageCell,
      size: 350,
    },
    {
      id: 'id',
      header: '',
      Cell: CommitHashCell,
    },
    {
      id: 'authorDate',
      header: '',
      Cell: CommitAuthorDateCell,
    },
    {
      id: 'detail',
      header: '',
      Cell: CommitDetailCell,
      size: 80,
      minSize: 80,
      maxSize: 80,
      grow: false,
    },
  ]

  return (
    <DataTable
      data={commits}
      columns={columns}
      pagination={pagination}
      page={page}
      loading={loading}
      emptyTitle={t('common.noResults')}
      onPageChange={onPageChange}
      tableOptions={{
        meta: {
          onDetailClick,
        },
      }}
    />
  )
}
