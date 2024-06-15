/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import { SystemItem } from "../applogic/model";

/* yankiebar */

export default function ReadSystem(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<SystemItem>(
    "magic-people.system",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const system = readResult.data;
  return (
    <div>
      {system && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{system.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{system.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Version</div>
            <div>{system.version}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{system.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{system.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{system.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{system.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{system.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
