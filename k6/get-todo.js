import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    {
      duration: "15s",
      target: 100
    },
    {
      duration: "30s",
      target: 100
    },
    {
      duration: "15s",
      target: 0
    }
  ],
  thresholds: {
    http_req_duration: ["p(99)<1000", "p(95)<700", "avg<500"]
  }
};

const rows = [
  { id: "e835b8ee-d8fe-42e7-9c84-99747c3ab675" },
  { id: "6b33fa6c-1d94-473c-84bf-37c205cb34fb" },
  { id: "f2249998-43df-4a58-bbb7-521aa697f7d5" },
  { id: "d48bdb5a-2e5b-40df-b4fc-949070699505" },
  { id: "22a7b34d-6528-49e8-aae7-32bce9c69c40" },
  { id: "99ebd1b8-c205-4814-8ef4-661be2b63715" },
  { id: "4de204ce-5f98-4c30-b4d3-5c2e2f7884e7" },
  { id: "a23dfd59-f630-44eb-9029-90ba15491e30" },
  { id: "63445e9c-5e90-4f03-9e4f-ae8b29e48ecd" },
  { id: "c924e4e3-16df-4fe9-bf26-2d5db8b7713b" },
  { id: "7dbe60e8-d2ec-4b27-a020-17c643c1a8b4" },
  { id: "677416a2-54ad-4bb3-adcc-1458cd18640a" },
  { id: "8125a86b-b2d0-4504-88a4-2898f6aaa7f0" },
  { id: "36743b02-658a-49af-99c1-49d83a4a7a73" },
  { id: "4d1a7eb3-ee51-4e66-86c8-d61f52ddc6d3" },
  { id: "d7781204-ee1d-4bfb-9636-d12527f5af10" },
  { id: "f61c1671-4ad2-4296-9f45-f5c6e7032cab" },
  { id: "1bc93542-9896-495f-8322-217014b8d022" },
  { id: "5befb4a9-b39e-451b-8a34-66653c10d598" },
  { id: "50352d04-cbd1-4292-a14f-d740c77b6b05" },
  { id: "f722aea8-2c84-4f98-a39b-30dcf6b3f784" },
  { id: "65f122de-ec61-45a7-815f-c9ed9a599584" },
  { id: "ec443f2e-a575-4967-8e03-cc3d8a56ba72" },
  { id: "5b6c70b4-eae6-4e60-bc2e-a9de8ec9430d" },
  { id: "b53336a6-97ae-4547-aeef-2c5f5d1844a3" },
  { id: "71856c05-d2a8-4a17-8bf8-02dc1582a4bf" },
  { id: "79084a40-1ee3-437e-9a15-501726699254" },
  { id: "72320d5d-4b44-40ab-9c24-4f3bd4c9b131" },
  { id: "5fc752c0-b057-40b4-9538-b8fa0b5515da" },
  { id: "dfe95754-c28a-430d-be05-3c34c46c0061" },
  { id: "6316a538-f3ed-4221-963e-5626020910c2" },
  { id: "eff42ba4-77d5-49ff-9aca-6233544a5a09" },
  { id: "a2817b16-88c3-453d-a940-e68b7c174d4e" },
  { id: "0fbf538b-2ea3-425c-9050-a8261e8c1c41" },
  { id: "ae348a4e-7df3-49dc-b97f-798d21854065" },
  { id: "4240894d-1657-46f3-8d89-b6134a4827f5" },
  { id: "cae573a8-2c88-4564-a715-4bb2569b6950" },
  { id: "3f612e1d-f5ba-4883-a2e7-60fac7156129" },
  { id: "e6accfc5-6656-4e86-9829-184d7614528f" },
  { id: "818f6711-2262-403d-bc90-54a75c4a2905" },
  { id: "ecf0e616-1bbb-4333-aab9-bfb4a20c6090" },
  { id: "3d5573c9-dde5-4be7-91c0-4c8596926b36" },
  { id: "daf952d2-38dc-4348-845b-d40bf025a990" },
  { id: "f7f5a47b-8efa-467b-b084-f63ccf5c9ffc" },
  { id: "3de390fb-a167-4fd4-917d-5e3bfed26c8d" },
  { id: "92846fde-2c6a-4b54-88b1-5fc5ba39ce4e" },
  { id: "d3be251b-47f7-430f-bbf1-d65e5cf4a54f" },
  { id: "c88ff3a3-0d0e-4d5c-8d5b-680e84922d0e" },
  { id: "c02aa55a-7ed8-4e34-a7b9-a3ee9c708f86" },
  { id: "c8bcfbac-cd35-4009-839f-600361c5c0c5" },
  { id: "f9b3bf8b-1a60-42bd-a867-d1c5ac8e3bce" },
  { id: "5e8e2d0e-2d3e-4bf8-b872-859ca102b28a" },
  { id: "b9e618ab-0214-4018-bc78-09c1f7dcc15c" },
  { id: "ad787670-4b09-4cc7-bb78-198da30ac8aa" },
  { id: "c25bb79b-dbca-4bf3-a917-80264cf86b7e" },
  { id: "0debb532-f01b-4f84-8163-b943984dfd53" },
  { id: "02270ffa-a01b-43a9-b02b-f397594ad557" },
  { id: "7e11c16e-2570-45d2-a93d-0bd94e4328d2" },
  { id: "87262a11-cd6a-4c60-9b9c-ceb3cc45dc47" },
  { id: "c5aa4b93-8669-432e-a05d-dc908cf5bebf" },
  { id: "eb8b377b-8cda-4a32-8471-d7cc8a3c1d5a" },
  { id: "ec19869f-4f34-4ee4-a2e9-35fc8712170f" },
  { id: "98a7220c-fa19-43c7-8303-2e6789c82f88" },
  { id: "18351663-d673-4ddb-b104-1a46d8ffeee0" },
  { id: "382a532e-1466-456a-ac61-8b7795600e2f" },
  { id: "13d009b4-7d31-450b-9766-7f0ed5e7936c" },
  { id: "c9dc0222-0ca9-4328-a5f3-2f4cbb45e124" },
  { id: "dc8a7ade-880d-4e12-a2ca-9dca72ea753a" },
  { id: "0ee18db5-ce69-49bc-ada7-ccc0e28e5c44" },
  { id: "ba729ff6-0a74-4933-baa9-41fdd4f742e6" },
  { id: "189974fe-ee0f-42cc-aa74-4f4fad5ee11c" },
  { id: "29ff655f-6e56-4761-b3eb-0ea2aec751e3" },
  { id: "5cd49e70-ecdf-4c5c-99a0-31ab7e0dba38" },
  { id: "5f3d4862-cfa7-4066-9e6e-0a33deea506f" },
  { id: "bf144b63-b565-4789-8ce8-763f77248d49" },
  { id: "8fa9be24-8984-48bf-b9be-8dc96dc034b7" },
  { id: "8e7e7950-677a-40a3-af0e-5972563708d4" },
  { id: "397c78bb-4a81-435b-9905-6dd4504c8a8e" },
  { id: "58c9ab34-4d8c-4f89-948f-c85e50c7ecd4" },
  { id: "91d7bd3d-17f7-4845-966a-b50cbe6a6393" },
  { id: "72174090-d990-4763-a502-c8b99602297f" },
  { id: "d3ffde32-dbba-4fd4-af9d-e175766e83e8" },
  { id: "a8a35279-4cf2-43df-9626-3f40927ffa23" },
  { id: "4d9a6a44-aea3-4ab5-ab01-320e9e2c78f1" },
  { id: "fbeb037f-7590-45c4-919b-84c69260f6ba" },
  { id: "d15a95c9-c947-4335-8a81-5a70e48b630f" },
  { id: "af8c42d1-d1c6-429a-856a-2de54cff6a16" },
  { id: "76f44cd1-306b-4af4-9acd-d167609948d8" },
  { id: "71ead06e-4991-4874-a33f-bea4625a66da" },
  { id: "b30a8324-f6c4-44b2-b0f8-90899225cd65" },
  { id: "8f1a0669-bbf4-44e8-8d44-1b69b925fca1" },
  { id: "e9b79397-a68f-4df3-aefb-c6dd7b02dd16" },
  { id: "866fae2f-9006-4c8a-bac9-66061936ab7c" },
  { id: "acd7cbe8-f8b5-4f20-9b14-72167d281e97" },
  { id: "7d2882d2-343d-4b6f-b2ba-9090dc821f1d" },
  { id: "0ab8e060-7d13-49f1-bf62-6554e7818979" },
  { id: "df2eee44-6279-4de5-af89-8261197a9287" },
  { id: "ecaf401a-a3b1-4686-8c6a-8080da07ecdb" },
  { id: "9ba8a0dc-e39c-42a7-b9cf-0c772fe5f832" },
  { id: "66ad436d-36ad-4be8-9164-92d4270c5c94" }
];

export default function () {
  const row = rows[Math.floor(Math.random() * rows.length)];
  const res = http.get(`http://localhost:8080/todo/${row.id}`);

  check(res, {
    "response code was 200": (res) => res.status === 200
  });

  sleep(2);
}
