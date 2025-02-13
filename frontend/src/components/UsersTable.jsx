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
  const [users, setUsers] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const openRoute = useNavigate();

  const { data, isLoading, error } = useQuery({
    queryFn: () => userService.getUsers(currentPage),
    queryKey: ["users", currentPage],
    enabled: !!currentPage,
  });

  useEffect(() => {
    setUsers(getUsers());
  }, []);

  function changePage(page) {
    setCurrentPage(page);
  }

  return (
    <section className="users-table w-full flex flex-col items-end">
      <main className="table-main w-full rounded-lg border border-solid mb-6">
        <div className="w-full flex items-center">
          <div className="header-cell one text-xs">Full Name</div>
          <div className="header-cell two text-xs">Email Address</div>
          <div className="header-cell three text-xs">Address</div>
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
              data?.data.map(({ id, name, email, address }) => (
                <div key={uuid()} className="detail-row flex items-center border-b cursor-pointer" onClick={() => {
                  openRoute(`${routeNames.posts}/${id}`);
                }}>
                  <div className="detail-cell one">
                    <p className="detail-txt user-name font-medium text-sm">{name}</p>
                  </div>
                  <div className="detail-cell two">
                    <p className="detail-txt user-email text-sm">{email}</p>
                  </div>
                  <div className="detail-cell three">
                    <p className="detail-txt user-address text-sm">{address}</p>
                  </div>
                </div>
              ))
        }
      </main>
      <div className="pagination-container">
      {
        data && data.totalPages && <PaginationComponent
          totalPages={data.totalPages}
          onPageChange={changePage}
          page={currentPage}
        />
      }
      </div>
    </section>
  )
}