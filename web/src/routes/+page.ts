import type { PageLoad } from './$types';

type GradeInfo = {
	student_id: string;
	grade: number | null;
};

type LatestData = {
	semester: string;
	name: string;
	info: GradeInfo[];
};

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/grade/latest`);
	const data: LatestData = await res.json();

	return { data };
};
