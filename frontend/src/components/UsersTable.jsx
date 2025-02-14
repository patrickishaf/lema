import "../styles/UsersTable.css";
import { useEffect, useState } from "react";
import { getUsers } from "../data/users";
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"
import { useNavigate } from "react-router-dom";
import routeNames from "@/navigation/routenames";
import { useQuery } from "@tanstack/react-query";
import userService from "@/services/user.service";
import Loader from "./Loader";
import uuid from "react-uuid";
import PaginationComponent from "./PaginationComponent";

export default function UsersTable() {
  const [currentPage, setCurrentPage] = useState(1);
  const [userList, setUserList] = useState([]);
  const openRoute = useNavigate();

  const { data: users, isLoading, error } = useQuery({
    queryFn: () => userService.getUsers(currentPage),
    queryKey: ["users", currentPage],
    enabled: !!currentPage,
  });

  function changePage(page) {
    setCurrentPage(page);
  }

  useEffect(() => {
    if (!users || !users.data || users.data.length === 0) return;

    setUserList(users.data)
  }, [users]);

  return (
    <section className="users-table w-full flex flex-col items-end">
      <div className="table-box w-full flex mb-6">
        <main className="table-main w-full rounded-lg border border-solid">
          <div className="w-full flex items-center">
            <div className="cell header one text-xs">Full Name</div>
            <div className="cell header two text-xs">Email Address</div>
            <div className="cell header three text-xs">Address</div>
          </div>
          {
            isLoading
              ?
                <Loader />
              : 
                error
                ?
                  <div>{error.message}</div>
                :
                userList.map(({ id, name, email, address }) => (
                  <div key={uuid()} className="detail-row flex items-center border-b cursor-pointer" onClick={() => {
                    openRoute(`${routeNames.posts}/${id}`);
                  }}>
                    <div className="cell detail one">
                      <p className="detail-txt user-name font-medium text-sm">{name}</p>
                    </div>
                    <div className="cell detail two">
                      <p className="detail-txt user-email text-sm">{email}</p>
                    </div>
                    <div className="cell detail three">
                      <p className="detail-txt user-address text-sm">{address}</p>
                    </div>
                  </div>
                ))
          }
        </main>
      </div>
      <div className="pagination-container">
      {
        users && users.totalPages && <PaginationComponent
          totalPages={users.totalPages}
          onPageChange={changePage}
          page={currentPage}
        />
      }
      </div>
    </section>
  )
}