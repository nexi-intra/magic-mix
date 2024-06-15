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
import { CountryItem } from "../applogic/model";

/* yankiebar */

export default function ReadCountry(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<CountryItem>(
    "magic-people.country",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const country = readResult.data;
  return (
    <div>
      {country && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{country.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{country.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">code</div>
            <div>{country.code}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{country.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{country.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{country.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{country.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{country.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
