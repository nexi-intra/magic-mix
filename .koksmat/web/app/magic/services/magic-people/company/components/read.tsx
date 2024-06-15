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
import { CompanyItem } from "../applogic/model";

/* yankiebar */

export default function ReadCompany(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<CompanyItem>(
    "magic-people.company",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const company = readResult.data;
  return (
    <div>
      {company && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{company.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{company.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Vat Number</div>
            <div>{company.vatnumber}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Phone Number</div>
            <div>{company.phonenumber}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Address</div>
            <div>{company.address}</div>
          </div>{" "}
          <div>
            <div className="font-bold">City</div>
            <div>{company.city}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Postal Code</div>
            <div>{company.postalcode}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Country</div>
            <div>{company.country_id}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{company.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{company.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{company.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{company.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{company.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
