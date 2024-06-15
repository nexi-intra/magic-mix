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
import { RelationItem } from "../applogic/model";

/* yankiebar */

export default function ReadRelation(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<RelationItem>(
    "magic-people.relation",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const relation = readResult.data;
  return (
    <div>
      {relation && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{relation.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{relation.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Company Relation</div>
            <div>{relation.relation_id}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Relation Type</div>
            <div>{relation.relationtype}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{relation.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{relation.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{relation.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{relation.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{relation.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
