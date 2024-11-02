import type { PageLoad } from './$types';

type GradeInfo = {
    student_id: string;
    grade: number;
};

export const load: PageLoad = async ({ fetch }) => {
    const t = new Date();
    const semester = `${t.getFullYear()}${t.getMonth() > 9 ? 'f' : 's'}`;

    const res = await fetch(`/api/grade/${semester}/HW0`);
    const info: GradeInfo[] = await res.json();

    return { info, semester };
};
