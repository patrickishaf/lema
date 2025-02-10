import "../styles/UsersPage.css";
import UsersTable from "./UsersTable";

export default function UsersPage() {
  return (
    <div className="users-page w-screen h-screen pt-32">
      <section className="w-full h-full">
        <h1 className="page-title text-6xl font-medium mb-6">Users</h1>
        <UsersTable />
      </section>
    </div>
  )
}