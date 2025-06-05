import * as React from "react";
import '../Index.css';
import {
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { ArrowUpDown, ChevronDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

export const columns = [
  {
    accessorKey: "Index",
    header: "Index",
    cell: ({ row }) => <div style={{background:"linear-gradient(#212121,#1c1c1c)",borderBottom:"5px",paddingLeft:"10px",paddingRight:"10px", borderRadius:"25%"}}>{row.getValue("Index")}</div>,
  },
  {
    accessorKey: "Timestamp",
    header: ({ column }) => (
      <Button
        variant="ghost"
        onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        className="btn"
      >
        Timestamp
        <ArrowUpDown className="ml-2 h-4 w-4" />
      </Button>
    ),
    cell: ({ row }) => <div className="text-muted-foreground">{row.getValue("Timestamp")}</div>,
  },
  {
    accessorKey: "Data",
    header: "Data",
    cell: ({ row }) => <Badge className="capitalize">{row.getValue("Data")}</Badge>,
  },
  {
    accessorKey: "PrevHash",
    header: () => <div className="text-right">Previous Hash</div>,
    cell: ({ row }) => (
      <div className="text-right text-xs break-words">{row.getValue("PrevHash")}</div>
    ),
  },
  {
    accessorKey: "Hash",
    header: () => <div className="text-right">Hash</div>,
    cell: ({ row }) => (
      <div className="text-right text-xs break-words">{row.getValue("Hash")}</div>
    ),
  },
];

export function BlockchainViewer() {
  const [blocks, setBlocks] = React.useState([]);
  const [sorting, setSorting] = React.useState([]);
  const [columnFilters, setColumnFilters] = React.useState([]);
  const [columnVisibility, setColumnVisibility] = React.useState({});
  const [rowSelection, setRowSelection] = React.useState({});
  const [isLoading, setIsLoading] = React.useState(true);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    const fetchBlocks = async () => {
      try {
        setIsLoading(true);
        const response = await fetch("http://localhost:8080/blocks");
        const data = await response.json();
        setBlocks(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchBlocks();
  }, []);

  const table = useReactTable({
    data: blocks,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  });

  return (
    <div className="block-container">
      <div className="w-full max-w-full">
        <h1 className="text-3xl font-bold mb-6 text-center">Go Blockchain Viewer</h1>
      </div>
        
      <div className="input-check">
        <div>
          <Input
            placeholder="Filter data..."
            value={table.getColumn("Data")?.getFilterValue() ?? ""}
            onChange={(event) =>
              table.getColumn("Data")?.setFilterValue(event.target.value)
            }
            className="max-w-sm"
          />
          </div>
          <div>
          <DropdownMenu >
            <DropdownMenuTrigger asChild>
              <Button variant="outline" className="btn">
                Columns <ChevronDown className="ml-2 h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="drp">
              <DropdownMenuLabel>Toggle Columns</DropdownMenuLabel>
              <DropdownMenuSeparator />
              {table
                .getAllColumns()
                .filter((column) => column.getCanHide())
                .map((column) => (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={column.getIsVisible()}
                    onCheckedChange={(value) =>
                      column.toggleVisibility(!!value)
                    }
                  >
                    {column.id}
                  </DropdownMenuCheckboxItem>
                ))}
            </DropdownMenuContent>
          </DropdownMenu>
        </div></div>
      <div>
        {isLoading && <p className="text-gray-500">Loading blocks...</p>}
        {error && <p className="text-red-500">Error: {error}</p>}

        {!isLoading && !error && (
          <div className="overflow-x-auto">
            <table className="min-w-full table-auto border border-gray-400 bg-white">
              <thead className="bg-gray-200">
                {table.getHeaderGroups().map((headerGroup) => (
                  <tr key={headerGroup.id}>
                    {headerGroup.headers.map((header) => (
                      <th
                        key={header.id}
                        // className=""
                      >
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </th>
                    ))}
                  </tr>
                ))}
              </thead>
              <tbody>
                {table.getRowModel().rows.map((row) => (
                  <tr
                    key={row.id}
                    className={row.getIsSelected() ? "bg-gray-100" : ""}
                  >
                    {row.getVisibleCells().map((cell) => (
                      <td
                        key={cell.id}
                        className="table-data"
                      >
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </td>
                    ))}
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        <div className="bottom-div">
          <div className="btn-div">
            <Button
            className="btn"
              variant="outline"
              size="xl"
              onClick={() => table.previousPage()}
              disabled={!table.getCanPreviousPage()}
            >
              Previous
            </Button>
            <Button
            className="btn"
              variant="outline"
              size="xl"
              onClick={() => table.nextPage()}
              disabled={!table.getCanNextPage()}
            >
              Next
            </Button>
          </div>
          <div className="text-sm text-gray-500">
            Page {table.getState().pagination.pageIndex + 1} of{" "}
            {table.getPageCount()}
          </div>
        </div>
      </div>
    </div>
  );
}

export default BlockchainViewer;
