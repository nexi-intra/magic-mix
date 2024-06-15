/**
 * v0 by Vercel.
 * @see https://v0.dev/t/poxPLZp338g
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
"use client";

import { useState, useMemo } from "react";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import { useSQLSelect } from "@/app/koksmat/usesqlselect";
import { CheckedState } from "@radix-ui/react-checkbox";

interface DataItem {
  name: string;
  id: string;
  count: number;
}

interface SortBy {
  key: keyof DataItem | null;
  order: "asc" | "desc" | null;
}

interface SQLResult {
  Result: DataItem[];
}
export function ListNameIdCount() {
  const importData = useSQLSelect<SQLResult>(
    "magic-mix.app",
    `
select name,row_number() OVER (ORDER BY name) AS id,count(*) 
	from importdata group by name order by count(*) desc
  `
  );

  const data = useMemo<DataItem[]>(() => {
    if (!importData.data) return [];
    return importData.data.Result;
  }, [importData.data]);

  const [selectedRows, setSelectedRows] = useState<DataItem[]>([]);
  const [sortBy, setSortBy] = useState<SortBy>({ key: null, order: null });

  const handleSort = (key: keyof DataItem) => {
    if (sortBy.key === key) {
      setSortBy({ key, order: sortBy.order === "asc" ? "desc" : "asc" });
    } else {
      setSortBy({ key, order: "asc" });
    }
  };

  const sortedData = useMemo(() => {
    if (!sortBy.key) return data;
    return [...data].sort((a, b) => {
      const sortOrder = sortBy.order === "asc" ? 1 : -1;
      if (a[sortBy.key!] < b[sortBy.key!]) return -1 * sortOrder;
      if (a[sortBy.key!] > b[sortBy.key!]) return 1 * sortOrder;
      return 0;
    });
  }, [data, sortBy]);

  const handleRowSelect = (item: DataItem, isSelected: boolean) => {
    setSelectedRows(
      isSelected
        ? [...selectedRows, item]
        : selectedRows.filter((row) => row !== item)
    );
  };

  const isRowSelected = (item: DataItem) => selectedRows.includes(item);

  return (
    <div className="flex flex-col gap-4">
      <div className="flex items-center gap-4">
        <Button variant="outline">
          <PencilIcon className="w-4 h-4 mr-2" />
          Edit
        </Button>
        <Button variant="outline">
          <TrashIcon className="w-4 h-4 mr-2" />
          Delete
        </Button>
        <span className="ml-auto text-sm">
          {selectedRows.length} row(s) selected
        </span>
      </div>
      <div className="overflow-auto border rounded-lg">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[40px]">
                <Checkbox
                  checked={selectedRows.length === data.length}
                  // indeterminate={
                  //   selectedRows.length > 0 && selectedRows.length < data.length
                  // }
                  onCheckedChange={() => {
                    if (selectedRows.length === data.length) {
                      setSelectedRows([]);
                    } else {
                      setSelectedRows(data);
                    }
                  }}
                />
              </TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => handleSort("name")}
              >
                Name
                {sortBy.key === "name" && (
                  <span className="ml-1">
                    {sortBy.order === "asc" ? "\u2191" : "\u2193"}
                  </span>
                )}
              </TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => handleSort("id")}
              >
                ID
                {sortBy.key === "id" && (
                  <span className="ml-1">
                    {sortBy.order === "asc" ? "\u2191" : "\u2193"}
                  </span>
                )}
              </TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => handleSort("count")}
              >
                Count
                {sortBy.key === "count" && (
                  <span className="ml-1">
                    {sortBy.order === "asc" ? "\u2191" : "\u2193"}
                  </span>
                )}
              </TableHead>
              <TableHead className="w-[40px]" />
            </TableRow>
          </TableHeader>
          <TableBody>
            {sortedData.map((item) => (
              <TableRow key={item.id}>
                <TableCell>
                  <Checkbox
                    checked={isRowSelected(item)}
                    onCheckedChange={(e: CheckedState) => {
                      if (typeof e === "boolean") handleRowSelect(item, e);
                    }}
                  />
                </TableCell>
                <TableCell className="font-medium">{item.name}</TableCell>
                <TableCell>{item.id}</TableCell>
                <TableCell>{item.count}</TableCell>
                <TableCell>
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="icon">
                        <MoveVerticalIcon className="w-4 h-4" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuItem>View</DropdownMenuItem>
                      <DropdownMenuItem>Edit</DropdownMenuItem>
                      <DropdownMenuItem>Delete</DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}

function MoveVerticalIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <polyline points="8 18 12 22 16 18" />
      <polyline points="8 6 12 2 16 6" />
      <line x1="12" x2="12" y1="2" y2="22" />
    </svg>
  );
}

function PencilIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z" />
      <path d="m15 5 4 4" />
    </svg>
  );
}

function TrashIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M3 6h18" />
      <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
      <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
    </svg>
  );
}
