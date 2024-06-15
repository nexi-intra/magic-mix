"use client";

import SearchConnection from "@/app/magic/services/magic-mix/connection/components/search";
import { DatasetEditOptions } from "@/components/dataset-edit-options";

export default function Page() {
  return (
    <div className="space-x-2 h-[90vh]">
      <DatasetEditOptions />
      <SearchConnection />
    </div>
  );
}
