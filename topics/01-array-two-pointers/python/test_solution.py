"""01 数组与双指针的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_rotate_array_shared_cases() -> None:
    for case in solution.load_shared_cases()["rotate"]:
        nums = list(case["nums"])
        solution.rotate_array(nums, case["k"])
        assert nums == case["expected"], case["name"]


def test_max_area_shared_cases() -> None:
    for case in solution.load_shared_cases()["max_area"]:
        assert solution.max_area(case["height"]) == case["expected"], case["name"]


def test_three_sum() -> None:
    assert solution.three_sum([-1, 0, 1, 2, -1, -4]) == [[-1, -1, 2], [-1, 0, 1]]
    assert solution.three_sum([0, 0, 0]) == [[0, 0, 0]]
    assert solution.three_sum([1, 2, 3]) == []
    assert solution.three_sum([0, 1]) == []


def test_move_zeroes() -> None:
    nums = [0, 1, 0, 3, 12]
    solution.move_zeroes(nums)
    assert nums == [1, 3, 12, 0, 0]


def test_trap() -> None:
    assert solution.trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]) == 6
    assert solution.trap([4, 2, 0, 3, 2, 5]) == 9
    assert solution.trap([]) == 0


def test_merge_intervals_does_not_mutate_input() -> None:
    intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
    assert solution.merge_intervals(intervals) == [[1, 6], [8, 10], [15, 18]]
    assert intervals == [[1, 3], [2, 6], [8, 10], [15, 18]]


def test_product_except_self() -> None:
    assert solution.product_except_self([1, 2, 3, 4]) == [24, 12, 8, 6]


def test_max_sub_array() -> None:
    assert solution.max_sub_array([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6
    assert solution.max_sub_array([-3, -1, -2]) == -1


def test_spiral_order() -> None:
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    assert solution.spiral_order(matrix) == [1, 2, 3, 6, 9, 8, 7, 4, 5]
    assert solution.spiral_order([]) == []


def test_set_zeroes() -> None:
    matrix = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]
    solution.set_zeroes(matrix)
    assert matrix == [[1, 0, 1], [0, 0, 0], [1, 0, 1]]


def test_rotate_matrix() -> None:
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    solution.rotate_matrix(matrix)
    assert matrix == [[7, 4, 1], [8, 5, 2], [9, 6, 3]]


def test_search_matrix() -> None:
    matrix = [
        [1, 4, 7, 11, 15],
        [2, 5, 8, 12, 19],
        [3, 6, 9, 16, 22],
        [10, 13, 14, 17, 24],
        [18, 21, 23, 26, 30],
    ]
    assert solution.search_matrix(matrix, 5) is True
    assert solution.search_matrix(matrix, 20) is False
