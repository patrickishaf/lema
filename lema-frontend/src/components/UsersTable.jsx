import "../styles/UsersTable.css";
import { useEffect, useState } from "react";
import { getUsers } from "../data/users";
import PaginationBtns from "./PaginationBtns";
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

export default function UsersTable() {
  const { data, isLoading, error } = useQuery({
    queryFn: () => userService.getUsers(),
    queryKey: ["users"]
  });
  const [users, setUsers] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const openRoute = useNavigate();

  useEffect(() => {
    setUsers(getUsers());
  }, []);

  function changePage(page) {
    const pageIndex = page - 1;
    setCurrentPage(pageIndex);
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
                  <p className="detail-cell user-name one font-medium text-sm">{name}</p>
                  <p className="detail-cell user-email two text-sm">{email}</p>
                  <p className="detail-cell user-address three text-sm">{address}</p>
                </div>
              ))
        }
      </main>
      <div className="pagination-container">
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <PaginationPrevious href="#" className="custom-pale-txt" />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink href="#" isActive className="custom-pale-txt">1</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationEllipsis />
            </PaginationItem>
            <PaginationItem>
              <PaginationNext href="#" className="custom-pale-txt" />
            </PaginationItem>
          </PaginationContent>
        </Pagination>
      </div>
    </section>
  )
}